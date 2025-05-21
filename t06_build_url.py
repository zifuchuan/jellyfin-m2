from base import app
from flask import url_for


@app.route("/test/<name>")
def test(name):
    return f"Hello, {name}!"


@app.route("/build_url")
def hello():
    return url_for('test', name='alice', age=11)
