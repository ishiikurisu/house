local basic_controller = require "house/controllers/basic_controller"
local JSON = require "dkjson"
local edit_controller = { }

edit_controller.construct = function(args)
  local self = basic_controller.new(args)

  -- Loading configuration file
  local configpath = '.'
  if self.repo ~= nil then
    configpath = './src/' .. self.repo
  end
  configpath = configpath .. '/.houseconfig'
  local config = util.readAll(configpath)
  if config ~= nil then
    self.params = JSON.decode(config, 1, nil).edit
  else
    self.params = self.options
  end

  return self
end

edit_controller.new = function(args)
  local self = edit_controller.construct(args)

  self.draw = function()
    local where = ' .'
    if self.repo ~= nil then
      where = ' src/' .. self.repo
    end
    local command = self.params.editor .. where .. ' &'
    local commands = { }
    table.insert(commands, command)
    self.execute(commands)
  end

  return self
end

return edit_controller
