from flask import Flask, request
from utils import *

app = Flask(__name__)


@app.route("/dsvw")
def dsvw():
    init()
    kind = request.args.get('kind', '')
    v = request.args.get('v', '')
    content = ''
    try:
        if kind == 'obj':
            content = query_obj(v)
        elif kind == 'id':
            content = query_id(v)
        elif kind == 'version':
            content = query_version(v)
        elif kind == 'path':
            content = query_path(v)
        elif kind == 'domain':
            content = query_domain(v)
        elif kind == 'xml':
            content = query_xml(v)
        elif kind == 'name':
            content = query_name(v)
        elif kind == 'include':
            content = query_include(v)
        elif kind == 'url':
            content = query_url(v)
        else:
            content = 'dsvw'
    except Exception as e:
        content = str(e)
    
    return content
