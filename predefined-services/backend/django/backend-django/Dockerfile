FROM python:${{DJANGO_PYTHON_VERSION}}-alpine

#? if services.database contains name == "postgres" {
#RUN apt-get install -y --no-install-recommends postgresql-client && rm -rf /var/lib/apt/lists/*  
#? }
WORKDIR /usr/src/app
COPY requirements.txt ./
RUN pip install -r requirements.txt
COPY . .

EXPOSE 8000
CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]