from flask import Flask

def create_app():
    app = Flask(__name__)
    app.config.from_pyfile("config.py")

    @app.route('/')
    def hello():
        return '<h1>Hello World</h1>'

    return app


def init_db():
    pass
    # from .models import db
    # Initialize db:
    # db.init_app(app)


if __name__ == "__main__":
    app = create_app()
    # init_db()
    app.run()