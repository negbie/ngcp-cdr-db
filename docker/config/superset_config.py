import os
from celery.schedules import crontab


def envvar(var_name, default=None):
    try:
        return os.environ[var_name]
    except KeyError:
        if default:
            return default
        raise EnvironmentError(f'The environment variable {var_name} is missing.')


POSTGRES_USER = envvar('POSTGRES_USER')
POSTGRES_PASSWORD = envvar('POSTGRES_PASSWORD')
POSTGRES_HOST = envvar('POSTGRES_HOST')
POSTGRES_PORT = envvar('POSTGRES_PORT')
POSTGRES_DB = envvar('POSTGRES_DB')

# The SQLAlchemy connection string.
SQLALCHEMY_DATABASE_URI = f'postgresql://{POSTGRES_USER}:{POSTGRES_PASSWORD}@{POSTGRES_HOST}:{POSTGRES_PORT}/{POSTGRES_DB}'

# SMTP configuration
EMAIL_NOTIFICATIONS = False  # all the emails are sent using dryrun
SMTP_HOST = 'localhost'
SMTP_STARTTLS = True
SMTP_SSL = False
SMTP_USER = 'superset'
SMTP_PORT = 25
SMTP_PASSWORD = 'superset'
SMTP_MAIL_FROM = 'superset@superset.com'

ENABLE_SCHEDULED_EMAIL_REPORTS = True
EMAIL_REPORT_FROM_ADDRESS = 'reports@superset.com'

REDIS_HOST = envvar('REDIS_HOST')
REDIS_PORT = envvar('REDIS_PORT')

CACHE_CONFIG = {
    'CACHE_TYPE': 'redis',
    'CACHE_DEFAULT_TIMEOUT': 60 * 60 * 24, # 1 day default (in secs)
    'CACHE_KEY_PREFIX': 'superset_results',
    'CACHE_REDIS_URL': f'redis://{REDIS_HOST}:{REDIS_PORT}/0',
}

class CeleryConfig(object):
    BROKER_URL = f'redis://{REDIS_HOST}:{REDIS_PORT}/0'
    CELERY_IMPORTS = (
        'superset.sql_lab',
        'superset.tasks',
    )
    CELERY_RESULT_BACKEND = f'redis://{REDIS_HOST}:{REDIS_PORT}/1'
    CELERYD_LOG_LEVEL = 'DEBUG'
    CELERYD_PREFETCH_MULTIPLIER = 1
    CELERY_ACKS_LATE = True
    CELERY_ANNOTATIONS = {
        'sql_lab.get_sql_results': {
            'rate_limit': '100/s',
        },
        'email_reports.send': {
            'rate_limit': '1/s',
            'time_limit': 120,
            'soft_time_limit': 150,
            'ignore_result': True,
        },
    }
    CELERY_TASK_PROTOCOL = 1
    CELERYBEAT_SCHEDULE = {
        'email_reports.schedule_hourly': {
            'task': 'email_reports.schedule_hourly',
            'schedule': crontab(minute=1, hour='*'),
        },
    }


CELERY_CONFIG = CeleryConfig

