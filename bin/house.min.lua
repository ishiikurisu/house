-- UTILITIES
local util = { }

function util.mysplit(inputstr, sep)
    if sep == nil then
        sep = "%s"
    end
    local t = { }; i = 1
    for str in string.gmatch(inputstr, "([^" .. sep .. "]+)") do
        t[i] = str
        i = i + 1
    end
    return t
end

-- BASIC CONTROLLER DEFINITION
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

-- LOAD CONTROLLER DEFINITION
local load_controller = { }

load_controller.new = function(args)
    local self = basic_controller.new(args)

    self.draw = function()
        local dirs = util.mysplit(self.args[1], "/")
        local levels = #dirs
        local commands = { }

        table.insert(commands, "cd src")
        for _, d in ipairs(dirs) do
            table.insert(commands, "cd " .. d)
        end
        table.insert(commands, "git pull origin master")
        for i = 1, levels do
            table.insert(commands, "cd ..")
        end

        self.execute(commands)
    end

    return self
end

-- UPLOAD CONTROLLER DEFINITION
local upload_controller = { }

upload_controller.new = function(args)
    local self = basic_controller.new(args)

    self.draw = function()
        local dirs = util.mysplit(self.args[1], "/")
        local levels = #dirs
        local commands = { }

        table.insert(commands, "cd src")
        for _, d in ipairs(dirs) do
            table.insert(commands, "cd " .. d)
        end
        table.insert(commands, "git add -A")
        table.insert(commands, "git commit")
        table.insert(commands, "git push origin master")
        for i = 1, levels do
            table.insert(commands, "cd ..")
        end

        self.execute(commands)
    end

    return self
end

-- GET CONTROLLER DEFINITION
local get_controller = { }

get_controller.new = function(args)
    local self = basic_controller.new(args)

    self.draw = function()
        local dirs = util.mysplit(self.args[1], "/")
        local levels = #dirs - 1
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

-- HOUSE DEFINITION
local house = { }

house.tools = {
    load = load_controller.new,
    upload = upload_controller.new,
    get = get_controller.new
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

-- MAIN PROCEDURE
print("---")
h = house.new(arg)
h.draw()
print("...")
