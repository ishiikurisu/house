local basic_controller = require "house/controllers/basic_controller"
local upload_controller = { }

upload_controller.new = function(args)
  local self = basic_controller.new(args)

  self.draw = function()
    local commands = { }
    local repo = self.args[1]

    if repo ~= nil then
      commands = self.buildTree(repo)
    else
      commands = self.addCommands(commands)
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
    commands = self.addCommands(commands)
    for i = 1, levels do
      table.insert(commands, "cd ..")
    end

    return commands
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
