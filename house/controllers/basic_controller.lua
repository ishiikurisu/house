local basic_controller = { }

basic_controller.construct = function(args)
    local self = { }

    self.args = args
    self.script_name = os.tmpname()
    self.shell = "sh"
    if self.script_name:sub(1, 1) == "\\" then
        self.shell = "cmd /C"
        self.script_name = "." .. self.script_name .. "bat"
    end

    return self
end

basic_controller.new = function(args)
    local self = basic_controller.construct(args)

    self.draw = function()
        print(self.args)
    end

    self.execute = function(commands)
        local fp = io.open(self.script_name, "w")
        io.output(fp)
        for _, cmd in ipairs(commands) do
            io.write(cmd .. "\n")
        end
        io.close(fp)
        os.execute(self.shell .. " " .. self.script_name)
        os.remove(self.script_name)
    end

    return self
end

return basic_controller
