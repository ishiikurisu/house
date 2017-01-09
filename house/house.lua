local house = { }

house.construct = function(args)
  local self = { }

  self.args = args

  return self
end

house.new = function(args)
  local self = house.construct(args)

  self.draw = function()
    if self.args == nil then
      print("nothing")
    else
      for i, it in ipairs(self.args) do
        print(i .. ". " .. it)
      end
    end
  end

  return self
end

return house
