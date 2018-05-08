class TokenValidationsController < DeviseTokenAuth::TokenValidationsController
  protected
  def render_validate_token_success
    render json: {
    success: true
  }
  end
end
