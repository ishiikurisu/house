local load_controller = require "house/controllers/load_controller"
local upload_controller = require "house/controllers/upload_controller"
local get_controller = require "house/controllers/get_controller"
local build_controller = require "house/controllers/build_controller"
local house = { }

house.tools = {
    load = load_controller.new,
    upload = upload_controller.new,
    get = get_controller.new,
    build = build_controller.new
}

house.construct = function(args)
    local self = { }
    local tool = args[1]

    table.remove(args, 1)
    self.tool = tool
    self.controller = house.tools[tool](args)

    return self
end

house.new = function(args)
    local self = house.construct(args)

    self.draw = function()
        self.controller.draw()
    end

    return self
end

return house
