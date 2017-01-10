local basic_controller = require "house/controllers/basic_controller"
local JSON = require "dkjson"
local build_controller = { }

build_controller.construc = function(args)
  local self = basic_controller.new(args)
  local configpath = 'src/' .. self.repo .. '/.houseconfig'

  -- Loading configuration file
  

  return self
end

build_controller.new = function(args)
  local self = build_controller.construct(args)

  self.draw = function()
    local commands = { }

    self.execute(commands)
  end

  return self
end

return build_controller
