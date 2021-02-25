Rails.application.routes.draw do
  resources :events do
    resources :comments
  end
end
