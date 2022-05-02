class AddEvent < ActiveRecord::Migration[6.1]
  def change
    add_column :comments, :event_id, :integer
  end
end
