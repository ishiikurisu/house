local flag_manager = { }

function flag_manager.get_available_flags()
  local flags = { }

  flags["-e"] = "editor"
  flags["-h"] = "help"

  return flags
end

function flag_manager.parse(args)
  local options = { }
  local key = "repo"
  local flags = flag_manager.get_available_flags()

  for i, arg in pairs(args) do
    if i > 0 then
      if flags[arg] ~= nil then
        key = flags[arg]
      else
        options[key] = arg
        key = "repo"
      end
    end
  end

  return options
end

return flag_manager
