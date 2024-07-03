class ProductConsumer < Karafka::BaseConsumer
  def consume(messages)
    messages.each do |message|
      product = message.payload
      ::Services::Api::V1::Products::Upsert.new(product_params.merge({is_api: false}), product).execute
    end
  end
end

private
  def product_params
    params&.permit(
      [
        :id,
        :name,
        :brand,
        :price,
        :description,
        :stock
      ]
    )
  end