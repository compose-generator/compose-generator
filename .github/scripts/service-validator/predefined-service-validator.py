from os.path import isfile, isdir, join
from os import listdir
from cerberus import Validator
import sys
import yaml
import json

def checkFileExistence():
    print('Checking file existence ...', end='')
    status = True
    template_path = '../../../predefined-services'
    template_types = [f for f in listdir(template_path) if isdir(join(template_path, f))]
    for template_type in template_types:
        template_type_path = template_path + '/' + template_type
        services = [f for f in listdir(template_type_path) if isdir(join(template_type_path, f))]
        for service in services:
            service_path = template_type_path + '/' + service
            # Check if service exists
            if not isfile(service_path + '/service.yml'):
                print('service.yml is missing for following service: ' + service)
                status = False
            # Check if config file exists
            if not isfile(service_path + '/config.json'):
                print('config.json is missing for following service: ' + service)
                status = False
            # Check if README exists
            if not isfile(service_path + '/README.md'):
                print('README.md is missing for following service: ' + service)
                status = False
    print(' done')
    return status

def checkYamlValidity():
    print('Checking YAML validity ...', end='')
    status = True
    schema = eval(open('./service-schema.py').read())
    v = Validator(schema)
    template_path = '../../../predefined-services'
    template_types = [f for f in listdir(template_path) if isdir(join(template_path, f))]
    for template_type in template_types:
        template_type_path = template_path + '/' + template_type
        services = [f for f in listdir(template_type_path) if isdir(join(template_type_path, f))]
        for service in services:
            service_path = template_type_path + '/' + service
            doc = loadYamlDoc(service_path + '/service.yml')
            if not v.validate(doc, schema):
                print('Invalid service file of stack: ' + service)
                status = False
    print(' done')
    return status

def checkJsonValidity():
    print('Checking JSON validity ...', end='')
    status = True
    schema = eval(open('./config-schema.py').read())
    v = Validator(schema)
    template_path = '../../../predefined-services'
    template_types = [f for f in listdir(template_path) if isdir(join(template_path, f))]
    for template_type in template_types:
        template_type_path = template_path + '/' + template_type
        services = [f for f in listdir(template_type_path) if isdir(join(template_type_path, f))]
        for service in services:
            service_path = template_type_path + '/' + service
            doc = loadJsonDoc(service_path + '/config.json')
            if not v.validate(doc, schema):
                print('Invalid config file of stack: ' + service)
                status = False
    print(' done')
    return status

def loadYamlDoc(path):
    with open(path, 'r') as stream:
        try:
            return yaml.safe_load(stream)
        except yaml.YAMLError as exception:
            raise exception

def loadJsonDoc(path):
    with open(path, 'r') as stream:
        try:
            return json.load(stream)
        except Exception as exception:
            raise exception

# Execute checks
if not checkFileExistence():
    sys.exit('File existence check failed.')
if not checkYamlValidity():
    sys.exit('Yaml validation check failed.')
if not checkJsonValidity():
    sys.exit('Json validation check failed.')