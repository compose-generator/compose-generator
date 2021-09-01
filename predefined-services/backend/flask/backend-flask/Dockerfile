FROM python:${{FLASK_PYTHON_VERSION}}-alpine
WORKDIR /usr/src/app

RUN pip install waitress

COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . .

CMD ["waitress-serve", "--port=5000", "--call", "run:create_app"]