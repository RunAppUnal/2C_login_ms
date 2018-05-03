
class RegistrationsController < DeviseTokenAuth::RegistrationsController
  protected
  def render_create_success
    render json: @resource
  end
end
