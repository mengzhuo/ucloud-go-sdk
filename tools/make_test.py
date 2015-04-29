#!/usr/bin/env python
# encoding: utf-8

import requests
import bs4
import sys
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

def make(url):
    struct_name = url.rpartition("/")[-1].split(".")[0].title().replace("_", '')
    data = requests.get(url).content
    p = bs4.BeautifulSoup(data)
    # Request
    

    return Template(tpl).substitute(sub_elem = sub_elem,
                                    struct_name=struct_name, 
                                    request_parameters=request_parameters, 
                                    response_parameters=response_parameters)

def search_for_url(url):
    last = url.rpartition("/")[0]
    node = bs4.BeautifulSoup(requests.get(url).content)
    url_list = []
    for a in node.select(".compound .internal"):
        url_list.append(last+"/"+a.get("href"))
    return url_list

if __name__ == '__main__':
    
    url = sys.argv[1]
    filepath = sys.argv[2]
    
    with open(filepath, "w+") as f:
        f.write("package ucloud\nimport testing\n")
        for u in search_for_url(url):
            print "Loading:" , u, 
            f.write(make(u).encode("utf8"))
            print "......Done"
