local basic_controller = require "house/controllers/basic_controller"
local JSON = require "dkjson"
local build_controller = { }

build_controller.construct = function(args)
  local self = basic_controller.new(args)

  -- Loading configuration file
  local configpath = '.'
  if self.repo ~= nil then
    configpath = 'src/' .. self.repo
  end
  configpath = configpath .. '/.houseconfig'
  local config = util.readAll(configpath)
  self.params = JSON.decode(config, 1, nil).build

  return self
end

build_controller.new = function(args)
  local self = build_controller.construct(args)

  self.draw = function()
    local commands = { }

    if (self.params['local'] == true) and (self.repo ~= nil) then
      local dirs = util.mysplit(self.repo, '/')
      table.insert(commands, 'cd src')
      for _, d in ipairs(dirs) do
        table.insert(commands, 'cd ' .. d)
      end
      table.insert(commands, self.params.command)
      for _ = 1, 1 + #dirs do
        table.insert(commands, 'cd ..')
      end
    else
      table.insert(commands, self.params.command)
    end

    self.execute(commands)
  end

  return self
end

return build_controller
