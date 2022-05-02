json.extract! comment, :id, :title, :description, :created_at, :updated_at
json.url event_comments_url(comment, format: :json)
