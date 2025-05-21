from base import app


@app.route("/hello", methods=['POST'])
def hello():
    return "Hello, X!"
