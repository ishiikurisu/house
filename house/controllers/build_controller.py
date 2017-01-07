from house.controllers.basic_controller import *
import json

class BuildController(BasicController):
    """docstring for LoadController."""
    def __init__(self, args):
        super().__init__(args)

        # Loading configuration file
        configpath = 'src/{0}/.houseconfig'.format(self.args[0])
        with open(configpath, 'r') as fp:
            content = fp.read()
        self.params = json.loads(content)['build']

    def draw(self):
        commands = []
        dirs = self.args[0].split('/')
        build_command = self.params['command']

        # Checking if build is local
        if 'local' in self.params:
            if self.params['local']:
                commands.append('cd src')
                for direc in dirs:
                    commands.append('cd ' + direc)
                commands.append(build_command)
                levels = 1 + len(dirs)
                for _ in range(levels):
                    commands.append('cd ..')
            else:
                commands.append(build_command)
        else:
            commands.append(build_command)

        # Executing commands
        self.execute(commands)
