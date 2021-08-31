{
    'label': {
        'type': 'string',
        'required': True
    },
    'preselected': {
        'type': 'string',
        'required': True
    },
    'demoAppInitCmd': {
        'type': 'list'
    },
    'serviceInitCmd': {
        'type': 'list'
    },
    'files': {
        'type': 'list',
        'schema': {
            'type': 'dict',
            'schema': {
                'path': {
                    'type': 'string',
                    'required': True
                },
                'type': {
                    'type': 'string',
                    'required': True,
                    'allowed': ['service', 'env', 'docs', 'config']
                }
            }
        }
    },
    'questions': {
        'type': 'list',
        'schema': {
            'type': 'dict',
            'schema': {
                'text': {
                    'type': 'string',
                    'required': True
                },
                'type': {
                    'type': 'integer',
                    'required': True
                    'allowed': [1, 2, 3]
                },
                'options': {
                    'type': 'list'
                },
                'defaultValue': {
                    'type': 'string'
                },
                'validator': {
                    'type': 'string'
                },
                'variable': {
                    'type': 'string',
                    'required': True
                },
                'advanced': {
                    'type': 'boolean'
                }
            }
        }
    },
    'volumes': {
        'type': 'list',
        'schema': {
            'type': 'dict',
            'schema': {
                'text': {
                    'type': 'string',
                    'required': True
                },
                'defaultValue': {
                    'type': 'string',
                    'required': True
                },
                'variable': {
                    'type': 'string',
                    'required': True
                },
                'advanced': {
                    'type': 'boolean'
                }
            }
        }
    },
    'secrets': {
        'type': 'list',
        'schema': {
            'type': 'dict',
            'schema': {
                'name': {
                    'type': 'string',
                    'required': True
                },
                'variable': {
                    'type': 'string',
                    'required': True
                },
                'length': {
                    'type': 'integer',
                    'required': True
                }
            }
        }
    }
}
