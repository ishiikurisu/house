import os
import subprocess

class BasicController:
    def __init__(self, args):
        self.args = args
        self.os_sep = '/' if os.name == 'posix' else '\\'
        self.call_file = 'ZXCVBNM.sh'

    def draw(self):
        print(self.args)

    def execute(self, commands):
        with open(self.call_file, 'w') as outlet:
            for command in commands:
                outlet.write('{0}\n'.format(command))
        subprocess.call(['sh', self.call_file])
        os.remove(self.call_file)
