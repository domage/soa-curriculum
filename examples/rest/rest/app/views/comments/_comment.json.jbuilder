json.extract! comment, :id, :title, :description, :created_at, :updated_at
json.url comment_url(comment, format: :json)
