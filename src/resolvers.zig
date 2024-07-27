const os = @import("std").os;

pub const CPU = struct {
    usage: f64,
    perCoreUsage: []f64,
    loadAverage: []f64,
};

pub const DISK = struct {
    total: u64,
    used: u64,
    free: u64,
    readBytes: u64,
    writeBytes: u64,
    readOperations: u64,
    writeOperations: u64,
};

pub const MEMORY = struct {
    total: u64,
    used: u64,
    free: u64,
    buffers: u64,
    cache: u64,
    swapTotal: u64,
    swapUsed: u64,
};

pub const NETWORK = struct {
    interface: []const u8,
    sentBytes: u64,
    receivedBytes: u64,
    sentPackets: u64,
    receivedPackets: u64,
    errors: u64,
    drops: u64,
};

pub const POWER = struct {
    status: []const u8,
    batteryLevel: f64,
};

pub const PROCESS = struct {
    name: []const u8,
    pid: u64,
    cpuUsage: f64,
    memoryUsage: f64,
};

pub const PROCESS = struct {
    name: []const u8,
    pid: u64,
    cpuUsage: f64,
    memoryUsage: f64,
};

pub const SYSTEM = struct {
    uptime: u64,
    cpu: CPU,
    memory: MEMORY,
    disks: []DISK,
    networks: []NETWORK,
    processes: []PROCESS,
    temperature: []f64,
    power: POWER
};

pub fn cpu() !CPU {
    // TODO: CPU usage

    // TODO: Per-core usage

    // TODO: Average usage

    var cpuMetrics = CPU{
        .usage = getCPUMetrics(),
        .perCoreUsage = getPerCoreUsage(),
        .loadAverage = getLoadAverage(),
    };
    return cpuMetrics;
}

// TODO: Define calls
