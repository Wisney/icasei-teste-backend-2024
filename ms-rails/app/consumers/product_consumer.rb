class ProductConsumer < Karafka::BaseConsumer
  def consume(messages)
    messages.each do |message|
      product = message.payload
      product['is_api'] = false
      product.symbolize_keys!
      
      ::Services::Api::V1::Products::Upsert.new(product, nil).execute
    end
  end
end