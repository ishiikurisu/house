from house.controllers.basic_controller import *

class LoadController(BasicController):
    """docstring for LoadController."""
    def __init__(self, args):
        super().__init__(args)

    def draw(self):
        commands = [
            'cd src',
            'echo $PWD',
            'cd ..'
        ]
        self.execute(commands)
