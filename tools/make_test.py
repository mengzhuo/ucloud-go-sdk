#!/usr/bin/env python
# encoding: utf-8

import requests
import bs4
import sys
import urlparse
from string import Template


tpl = u"""
// ---------------- Test${struct_name} ------------------
func Test${struct_name}(t *testing.T) {
    r := &${struct_name}{${mandatory_params}     
        }
    cmp := `${cmp_uri}`

    if err := FakeGetAndCmp(r, cmp); err != nil {
        t.Fatal(err)
    }
}

"""

def make(url, struct_name):
    data = requests.get(url).content
    p = bs4.BeautifulSoup(data)
    # Request
    cmp_uri = p.find("div", {"id":"request-example"}).find("pre").text.replace("\n", "").replace("(s)", "s")
    
    mandatory_params = ""
    params = urlparse.parse_qs(cmp_uri)
    for k, v in params.items():
        if k != "https://api.ucloud.cn/?Action":
            vv = v[0]
            if not vv.isdigit():
                vv = '"%s"' % vv

            print k,":=", vv
            mandatory_params += '%s:%s,\n' % (k, vv )

    return Template(tpl).substitute(mandatory_params= mandatory_params,
                                    struct_name=struct_name, 
                                    cmp_uri=cmp_uri)

def search_for_url(url):
    last = url.rpartition("/")[0]
    node = bs4.BeautifulSoup(requests.get(url).content)
    url_list = []
    for a in node.select(".compound .internal"):
        url_list.append((last+"/"+a.get("href"), a.text))
    return url_list

if __name__ == '__main__':
    
    url = sys.argv[1]
    filepath = sys.argv[2]
    
    with open(filepath, "w+") as f:
        f.write('package ucloud\nimport ("testing")\n')
        for u,s in search_for_url(url):
            print "Loading:" ,s, u, 
            f.write(make(u, s).encode("utf8"))
            print "......Done"
