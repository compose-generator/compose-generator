FROM python:${{DJANGO_PYTHON_VERSION}}-alpine

#? if services.database contains name == "postgres" {
#RUN apk update && apk add --no-cache postgresql-client && rm -rf /var/cache/apk/*
#? }
WORKDIR /usr/src/app
COPY requirements.txt ./
RUN pip install -r requirements.txt
COPY . .

EXPOSE 8000
CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]