from base import app


@app.route("/<path:path>")
def hello(path):
    return f"Hello, {path}!"
