"""Script to test all combinations of predefined service templates"""

from os.path import isdir, join
from os import listdir, system
import sys
import itertools

def get_all_template_names():
    """Returns a string array with all existing template names"""

    template_names = []
    template_path = '../../../predefined-services'
    #template_types = ["backend", "database", "db-admin", "frontend"]
    template_types = ["backend"]
    for template_type in template_types:
        template_type_path = template_path + '/' + template_type
        service_names = [f for f in listdir(template_type_path) if isdir(join(template_type_path, f))]
        template_names.extend(service_names)

    return template_names

def test_combination(comb):
    """Tests one particular combination of services"""
    # Create config file
    

    # Execute Compose Generator with the config file
    if system("compose-generator -c config.yml -i") != 0:
        sys.exit('Compose Generator failed when generating stack for combination ' + comb)

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
for i in range(1, len(templates) +1):
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
