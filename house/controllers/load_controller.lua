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
      commands = self.buildTree(repo)
    else
      table.insert(commands, self.command)
    end

    self.execute(commands)
  end

  self.buildTree = function(repo)
    local dirs = util.mysplit(repo, "/")
    local levels = #dirs
    local commands = { }

    table.insert(commands, "cd src")
    for _, d in ipairs(dirs) do
      table.insert(commands, "cd " .. d)
    end
    table.insert(commands, self.command)
    for i = 1, levels do
      table.insert(commands, "cd ..")
    end

      return commands
    end

    return self
end

return load_controller
