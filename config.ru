require 'erb'
require 'socket'

def erb(template)
  path = File.expand_path("../#{template}", __FILE__)
  ERB.new(File.read(path)).result(binding)
end

run Proc.new {|env| [200, {"Content-Type" => "text/html"}, [erb('index.erb')]]}
