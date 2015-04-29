#!/usr/bin/env python
# encoding: utf-8

import requests
import bs4
import sys
from string import Template


tpl = u"""
// ---------------- $struct_name ------------------
$sub_elem
type ${struct_name}Response struct {
	BaseResponse
        $response_parameters
}
func (r *${struct_name}Response) Data() interface{} {
	return
}

type $struct_name struct {
    $request_parameters
}

func (r *$struct_name) R() UResponse {
	return &${struct_name}Response{}
}

"""


TYPE_MAP = {u"String":u"string", u"Integer":u"int", u"Array":u"[]string"}
def make(url):
    struct_name = url.rpartition("/")[-1].split(".")[0].title().replace("_", '')
    data = requests.get(url).content
    p = bs4.BeautifulSoup(data)
    # Request
    req_t = p.select("#request-parameters table tbody tr")
    request_parameters = u""
    for row in req_t:
        tds = [x.text for x in row.select("td")]
        name = tds[0]

        type = TYPE_MAP.get(tds[1])
        if u"." in name:
            type = u"[]string"
            name = name.split(u".")[0]
        optional = u"" 
        try:
            if tds[3].lower() == "no":
                optional = u'`ucloud:"optional"`'
        except:
            pass
        doc = u"// " + tds[2].replace(u"\n", u"")
        request_parameters += u"{0} {1} {2} {3}\n".format(name, type, optional, doc)

    response_table = p.select("#response-elements table")[0].select("tbody tr")
    response_parameters = u""
    for row in response_table:
        tds = [x.text for x in row.select("td")]
        name = tds[0]
        if name in (u"Action", u"RetCode"):
            continue
        type = {u"String":u"string", u"Integer":u"int", u"Float":u"Float64"}.get(tds[1], "WARNING")
        doc = u"// " + tds[2]
        response_parameters += u'{0} {1} `json:",omitempty"`{2}\n'.format(name, type,  doc)
    
    sub_elem = u""
    if len(p.select("#response-elements table")) > 1:
        for i, t in enumerate(p.select("#response-elements table")[1:]):
            sub_struct = p.select("#response-elements strong")[i].text
            sub_elem += u"type {0} struct {{\n".format(sub_struct)
            for tds in [x.select("td") for x in t.select("tbody tr")]:
                name = tds[0].text
                type = {u"String":u"string", u"Integer":u"int", u"Float":u"float64"}.get(tds[1].text, "WARNING")
                doc = u"// " + tds[2].text.replace(u"\n", u"")
                sub_elem += u"{0} {1} {2}\n".format(name, type, doc)
            sub_elem += "}\n"

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
        f.write("package ucloud\n")
        for u in search_for_url(url):
            print "Loading:" , u, 
            f.write(make(u).encode("utf8"))
            print "......Done"
