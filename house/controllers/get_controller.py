from house.controllers.basic_controller import *

class GetController(BasicController):
    """docstring for LoadController."""
    def __init__(self, args):
        super().__init__(args)

    def draw(self):
        dirs = self.args[0].split('/')
        levels = 1 + len(dirs)
        commands = []

        commands.append('cd src')
        for direc in dirs:
            commands.append('cd ' + direc)
        commands.append('git clone https://{0}.git'.format(self.args[0]))
        for _ in range(levels):
            commands.append('cd ..')

        self.execute(commands)
