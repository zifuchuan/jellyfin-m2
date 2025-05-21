from base import app
from flask import render_template


@app.route('/hello/<name>')
def hello(name=None):
    return render_template('hello.html', name=name)
