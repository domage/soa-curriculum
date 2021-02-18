require 'sinatra'
require 'json'

events = [
    {
		id:          "1",
		title:       "Introduction to Golang",
		description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
	{
		id:          "2",
		title:       "Title number 2",
		description: "The second one",
	},
	{
		id:          "3",
		title:       "Title 3",
		description: "Desco",
	}
]

post '/event' do
    request.body.rewind
    event = JSON.parse request.body.read

    events << event

    JSON.dump events
end

delete '/events/:id' do
    events = events.select{ |event| event[:id] != params["id"] }
    ""
end

get '/events' do
  JSON.dump events
end

get '/events/:id' do
    JSON.dump events.select{ |event| event[:id] == params["id"] }.first
end
