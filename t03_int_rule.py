from base import app


@app.route("/<int:number>")
def hello(number):
    return f"Hello, {number}!"
