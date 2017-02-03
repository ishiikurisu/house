local basic_controller = require "house/controllers/basic_controller"
local JSON = require "dkjson"
local edit_controller = { }

edit_controller.construct = function(args)
  local self = basic_controller.new(args)

  -- Loading configuration file
  local configpath = './src/' .. self.repo .. '/.houseconfig'
  local config = util.readAll(configpath)
  self.params = JSON.decode(config, 1, nil).edit

  return self
end

edit_controller.new = function(args)
  local self = edit_controller.construct(args)

  self.draw = function()
    local where = ' src/' .. self.repo
    local command = self.params.editor .. where .. ' &'
    local commands = { }
    table.insert(commands, command)
    self.execute(commands)
  end

  return self
end

return edit_controller
