local basic_controller = require "house/controllers/basic_controller"
local upload_controller = { }

upload_controller.new = function(args)
  local self = basic_controller.new(args)

  self.draw = function()
    local commands = { }
    local repo = self.args[1]

    if repo ~= nil then
      commands = self.buildTree(repo, self.addCommands)
    else
      commands = self.addCommands(commands)
      table.insert(commands, #commands-2, "git checkout " .. self.script_name)
    end

    self.execute(commands)
  end

  self.addCommands = function(commands)
    table.insert(commands, "git add -A")
    table.insert(commands, "git commit")
    table.insert(commands, "git push origin master")
    return commands
  end

  return self
end

return upload_controller
