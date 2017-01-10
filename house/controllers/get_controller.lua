local basic_controller = require "house/controllers/basic_controller"
local get_controller = { }

get_controller.new = function(args)
    local self = basic_controller.new(args)

    self.draw = function()
        local dirs = util.mysplit(self.args[1], "/")
        local levels = #dirs
        local commands = { }

        table.insert(commands, "cd src")
        for _, d in ipairs(dirs) do
            table.insert(commands, "cd " .. d)
        end
        table.insert(commands, "git clone https://" .. self.args[1] .. ".git")
        for i = 1, levels do
            table.insert(commands, "cd ..")
        end

        self.execute(commands)
    end

    return self
end

return get_controller
