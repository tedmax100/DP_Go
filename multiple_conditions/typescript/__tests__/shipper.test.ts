describe("shipper cart test cases", () => {
  let blackCat: string;
  let hsinChu :string;
  let postOffice: string;
  let cart: Cart;

  beforeAll(() =>{
    blackCat = "black cat";
    hsinChu = "hsinChu";
    postOffice = "post office";
  })

  beforeEach(() => {
    cart = new Cart();
  })
  it("black cat with light weight", ()=> {
    let shippingFee = cart.shippingFee(blackCat, 30,20,10,5);
    expect(150).toEqual(shippingFee);
  })

  it("black cat with heavy weight", ()=> {
    let shippingFee = cart.shippingFee(blackCat, 30, 20, 10, 50);
    expect(500).toEqual(shippingFee);
  })
  
  it("hsinChu with small size", ()=> {
    let shippingFee = cart.shippingFee(hsinChu, 30,20,10,50);
    expect(144).toEqual(shippingFee);
  })

  it("hsinChu with huge size", ()=> {
    let shippingFee = cart.shippingFee(hsinChu, 110,20,10,50);
    expect(984).toEqual(shippingFee);
  }) 

  it("post office by weight", ()=> {
    let shippingFee = cart.shippingFee(postOffice, 100, 20, 10, 3);
    expect(110).toEqual(shippingFee);
  })

  it("post office by size", ()=> {
    let shippingFee = cart.shippingFee(postOffice, 100, 20, 10, 300);
    expect(440).toEqual(shippingFee);
  }) 

  it("unknown shipper", ()=> {
    expect( () => {
      cart.shippingFee("sinyi", 100, 20, 10, 300)
    }).toThrow("shipper not exist")  
  }) 
})

class Cart {
  shippingFee = (shipper: string, length : number, width : number, height : number, weight : number) : number => {
    if (shipper === "black cat") {
        if (weight > 20) {
            return 500;
        } else {
            return 100 + weight * 10;
        }
    } else if (shipper === "hsinChu") {
        let size: number;
        size = length * width * height;
        if (length > 100 || width > 100 || height > 100) {
            return size * 0.00002 * 1100 + 500;
        } else {
            return size * 0.00002 * 1200;
        }
    } else if (shipper === "post office") {
        let feeByWeight, size, feeBySize: number;
        feeByWeight = 80 + weight * 10;
        size = length * width * height;
        feeBySize = size * 0.00002 * 1100;
        return feeByWeight < feeBySize ? feeByWeight : feeBySize;
    } else {
        throw new Error("shipper not exist");

    }
  }
}