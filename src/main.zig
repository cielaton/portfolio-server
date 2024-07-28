const std = @import("std");
const zap = @import("zap");

// Create a hash map with string keys and request function values
var routes: std.StringHashMap(zap.HttpRequestFn) = undefined;

// Match the request routes
fn on_request(r: zap.Request) void {
    if (r.path) |thePath| {
        if (routes.get(thePath)) |callback| {
            callback(r);
            return;
        }
    }

    // Default
    r.sendJson("error: \"No matched route\"") catch return;
}

pub fn main() !void {
    var listener = zap.HttpListener.init(.{ .port = 3000, .on_request = on_request, .log = true });
    try listener.listen();

    std.debug.print("Listening on 0.0.0.0:3000\n", .{});

    zap.start(.{ .threads = 2, .workers = 3 });
}
