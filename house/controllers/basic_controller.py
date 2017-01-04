import os
import subprocess

class BasicController:
    def __init__(self, args):
        self.args = args
        self.os_sep = '/' if os.name == 'posix' else '\\'
        self.shell = 'sh' if os.name == 'posix' else 'cmd'
        self.call_file = 'ZXCVBNM.sh'

    def draw(self):
        print(self.args)

    def execute(self, commands):
        commands.append('exit')
        stack = [self.shell]
        with open(self.call_file, 'w') as outlet:
            for command in commands:
                outlet.write('{0}\n'.format(command))
        if self.shell == 'cmd':
            stack.append('/C')
        stack.append(self.call_file)
        subprocess.call(stack)
        os.remove(self.call_file)
