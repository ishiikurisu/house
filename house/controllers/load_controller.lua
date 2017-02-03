local basic_controller = require "house/controllers/basic_controller"
local load_controller = { }

load_controller.construct = function(args)
  local self = basic_controller.new(args)

  self.command = "git pull origin master"

  return self
end

load_controller.new = function(args)
  local self = load_controller.construct(args)

  self.draw = function()
    local commands = { }
    local repo = self.args[1]

    if repo ~= nil then
      commands = self.buildTree(repo, self.addCommands)
    else
      table.insert(commands, self.command)
    end

    self.execute(commands)
  end

  self.addCommands = function(commands)
    table.insert(commands, self.command)
    return commands
  end

  return self
end

return load_controller
