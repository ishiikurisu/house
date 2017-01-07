from house import controllers

class House:
    controllers2actions = {
        'load': controllers.LoadController,
        'upload': controllers.UploadController,
        'get': controllers.GetController
    }

    """Creates a house object with the appropriate controller depending
    on the given parameters. """
    def __init__(self, argv):
        # Extracting action and assigning correct controller
        self.action = argv[0]
        self.controller = self.controllers2actions[self.action](argv[1:])

    def draw(self):
        self.controller.draw()
