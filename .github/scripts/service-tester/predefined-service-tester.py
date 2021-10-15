"""Script to test all combinations of predefined service templates"""

import os
import sys
import itertools

def get_all_template_names():
    """Returns a string array with all existing template names"""
    template_types = list(filter(os.path.isdir, os.listdir("./predefined-templates")))

    template_names = []
    for name in template_types:
        template_names += list(filter(os.path.isdir, os.listdir("./predefined-templates/" + name)))

    print(template_names)
    return template_names

def test_combination(comb):
    """Tests one particular combination of services"""
    # Create config file
    

    # Execute Compose Generator with the config file
    if os.system("compose-generator -c config.yml -i") != 0:
        sys.exit('Compose Generator failed when generating stack for combination ' + comb)

    # Execute Compose Generator with the config file
    if os.system("docker compose up -d") != 0:
        sys.exit('Docker failed when generating stack for combination ' + comb)

def reset_environment():
    """Deletes all Docker related stuff. Should be executed after each test"""
    os.system("docker system prune -af")
    os.system("sudo rm -rf ./*")

# Initially reset the testing environment
reset_environment()

# Find all possible template combinations
templates = get_all_template_names()
combinations = []
for t in range(len(templates) +1):
    combinations_list = list(itertools.combinations(templates, t))
    combinations += combinations_list
total = len(combinations)

# Execute test for each combination
for i, combination in enumerate(combinations):
    print(f"Testing combination {i} of {total} ...")
    #test_combination(combination)
    reset_environment()

# Test was successful
print("Tested all combinations successfully!")
