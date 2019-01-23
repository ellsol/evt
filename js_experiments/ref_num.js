
const ByteBuffer = require("bytebuffer");


const last_irreversible_block_id = "01c2f8c6943e0295e09c30f7d7b80860a9f87c8447b90b85aad4375c18bb9813";
const hash = ByteBuffer.fromHex(last_irreversible_block_id, true); // little endian

const numHex = last_irreversible_block_id.substr(4, 4);
console.log(numHex)
const last_irreversible_block_num = ByteBuffer.fromHex(numHex, false).readUint16(0);


const last_irreversible_block_prefix = hash.readUInt32(8);


console.log("BlockNum: " + last_irreversible_block_num);
console.log("BlockPre:" + last_irreversible_block_prefix);