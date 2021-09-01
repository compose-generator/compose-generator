FROM python:${{TELEGRAM_BOT_PYTHON_VERSION}}-alpine

WORKDIR /usr/src/app
COPY requirements.txt ./
RUN pip install -r requirements.txt
COPY . .

CMD ["python", "bot.py"]