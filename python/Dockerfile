FROM python:3.12-alpine AS builder

RUN pip install --no-cache-dir poetry

WORKDIR /app
ENV POETRY_VIRTUALENVS_CREATE=false

COPY pyproject.toml poetry.lock ./

RUN poetry config virtualenvs.create false && poetry install --only main --no-interaction --no-ansi --no-root

FROM python:3.12-alpine AS production

WORKDIR /app

COPY --from=builder /usr/local/lib/python3.12 /usr/local/lib/python3.12
COPY --from=builder /usr/local/bin /usr/local/bin

COPY manage.py ./
COPY onxp/ ./onxp/

ENV PYTHONOPTIMIZE=1
ENV PYTHONDONTWRITEBYTECODE=1

EXPOSE 8383

CMD ["python", "manage.py", "runserver", "0.0.0.0:8383"]
