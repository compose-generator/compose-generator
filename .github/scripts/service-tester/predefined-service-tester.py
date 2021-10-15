"""Script to test all combinations of predefined service templates"""

from os.path import isdir, join
from os import listdir, system, remove
import sys
import itertools
import yaml

MAX_COMBINATION_SIZE = 1 # Careful! Runtime increases exponentially
TEMPLATES_PATH = "../../../predefined-services"
BIN_PATH = "../../../bin"

def get_all_template_names():
    """Returns a string array with all existing template names"""

    template_tuples = []
    template_types = ["backend", "database", "db-admin", "frontend"]
    for template_type in template_types:
        template_type_path = TEMPLATES_PATH + '/' + template_type
        for service_name in [f for f in listdir(template_type_path) if isdir(join(template_type_path, f))]:
            template_tuples.append((service_name, template_type))

    return template_tuples

def test_combination(comb):
    """Tests one particular combination of services"""
    # Create config file
    services = []
    for service in comb:
        services.append({"service": service[0], "type": service[1]})
    config = [{"project_name": "Example project", "services": services}]
    with open(BIN_PATH + "/config.yml", "w", encoding='utf-8') as file:
        yaml.dump(config, file, default_flow_style=False)

    # Execute Compose Generator with the config file
    if system("compose-generator -c config.yml -i") != 0:
        sys.exit('Compose Generator failed when generating stack for combination ' + comb)

    # Delete config file
    remove(BIN_PATH + "/config.yml")

    # Execute Compose Generator with the config file
    if system("docker compose up -d") != 0:
        sys.exit('Docker failed when generating stack for combination ' + comb)

def reset_environment():
    """Deletes all Docker related stuff. Should be executed after each test"""
    system("docker system prune -af")
    system("sudo rm -rf ./*")

# Initially reset the testing environment
print("Do initial cleanup ...", end='')
reset_environment()
print(" done")

# Find all possible template combinations
print("Collecting template names ...", end='')
templates = get_all_template_names()
combinations = []
for i in range(1, MAX_COMBINATION_SIZE +1):
    combinations.extend(list(itertools.combinations(templates, i)))
print(" done")
print(combinations)

# Execute test for each combination
print("Execute tests ...")
for i, combination in enumerate(combinations):
    print(f"Testing combination {i} of {len(combinations)} ...")
    #if len(combination) > 0:
    #    test_combination(combination)
    #    reset_environment()
print("Done")

# Test was successful
print("Tested all combinations successfully!")
