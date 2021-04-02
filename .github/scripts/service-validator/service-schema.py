{
    'build': {
        'type': 'string'
    },
    'image': {
        'type': 'string'
    },
    'container_name': {
        'type': 'string',
        'required': True
    },
    'volumes': {
         'type': 'list'
    },
    'networks': {
         'type': 'list',
         'nullable': True
    },
    'ports': {
         'type': 'list'
    },
    'env_file': {
        'type': 'list'
    },
    'environment': {
        'type': 'list'
    },
    'restart': {
        'type': 'string',
        'allowed': ['no', 'always', 'on-failure', 'unless-stopped']
    },
    'depends_on': {
        'type': 'list',
         'nullable': True
    },
    'links': {
        'type': 'list',
         'nullable': True
    },
    'labels': {
        'type': 'list',
         'nullable': True
    },
    'command': {
        'type': 'string'
    },
}