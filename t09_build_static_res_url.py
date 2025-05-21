from base import app
from flask import url_for


@app.route("/hello")
def hello():
    return url_for('static', filename='style.css')
