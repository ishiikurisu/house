local basic_controller = require "house/controllers/basic_controller"
local load_controller = { }

load_controller.new = function(args)
    local self = basic_controller.new(args)

    self.draw = function()
        self.execute(self.args)
    end

    return self
end

return load_controller
