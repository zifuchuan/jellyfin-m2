from base import app


@app.route("/about/")
def hello():
    return f"Hello, X!"
