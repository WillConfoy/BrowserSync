FROM python:3.12-alpine AS server

WORKDIR /usr/src/flask

COPY requirements.txt requirements.txt

RUN pip install --no-cache-dir -r requirements.txt

COPY . /usr/src/flask/

EXPOSE 5000

CMD [ "flask", "run" ]